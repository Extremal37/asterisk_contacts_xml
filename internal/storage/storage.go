package storage

import (
	"context"
	"database/sql"
	"github.com/Extremal37/asterisk_contacts_xml/internal/entity"
	"github.com/Extremal37/asterisk_contacts_xml/internal/logger"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const connTimeout = 3 * time.Minute

type Storage struct {
	conn *sql.DB
	log  *logger.Logger
}

// New receive sql connection and logger and return Storage
func New(conn *sql.DB, log *logger.Logger) *Storage {
	return &Storage{
		conn: conn,
		log:  log,
	}
}

func (s *Storage) Contacts(ctx context.Context) ([]entity.ContactDB, error) {
	// Create context with timeout defined in connTimeout
	c, cancel := context.WithTimeout(ctx, connTimeout)
	defer cancel()

	//Prepare query
	var contacts []entity.ContactDB
	query := "SELECT name,callerid FROM subscribers"

	// Fetching rows from DB
	rows, err := s.conn.QueryContext(c, query)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			s.log.Errorf("Unable to close rows : %v", err)
		}
	}()

	// Scan rows and put data to contacts array
	for rows.Next() {
		row := entity.ContactDB{}
		err = rows.Scan(&row.Number, &row.Name)
		if err != nil {
			s.log.Warnf("Unable to scan row: %v", err)
			continue
		}
		contacts = append(contacts, row)
	}

	return contacts, nil
}

func (s *Storage) Close() {
	err := s.conn.Close()
	if err != nil {
		s.log.Warnf("Failed to close DB connection: %v", err)
	} else {
		s.log.Debug("DB connection closed")
	}
}
