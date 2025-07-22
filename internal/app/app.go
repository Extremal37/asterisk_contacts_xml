package app

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"github.com/Extremal37/asterisk_contacts_xml/internal/config"
	"github.com/Extremal37/asterisk_contacts_xml/internal/entity"
	"github.com/Extremal37/asterisk_contacts_xml/internal/logger"
	"github.com/Extremal37/asterisk_contacts_xml/internal/storage"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const xmlName = "contact"

type App struct {
	config  *config.Config
	log     *logger.Logger
	storage *storage.Storage
}

func New(config *config.Config, log *logger.Logger) *App {
	return &App{config: config, log: log}
}

func (a *App) Run(ctx context.Context) error {
	a.log.Info("Initializing DB connection...")
	conn, err := initDB(a.config)
	if err != nil {
		return fmt.Errorf("failed to init database: %s", err)
	}

	a.storage = storage.New(conn, a.log)
	defer a.storage.Close()

	contacts, err := a.contacts(ctx)
	if err != nil {
		return fmt.Errorf("failed to get contacts: %s", err)
	}

	err = a.exportXML(contacts)
	if err != nil {
		return fmt.Errorf("failed to export XML: %s", err)
	}

	return nil
}

func initDB(cfg *config.Config) (*sql.DB, error) {
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database))
	if err != nil {
		return conn, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = conn.Ping()
	if err != nil {
		return conn, fmt.Errorf("failed to ping database: %w", err)
	}

	return conn, nil
}

func (a *App) contacts(ctx context.Context) ([]entity.ContactDB, error) {
	return a.storage.Contacts(ctx)

}

func (a *App) exportXML(contacts []entity.ContactDB) error {
	contactsXML := make([]entity.ContactXML, len(contacts))
	a.log.Infof("Exporting %d contacts", len(contacts))
	for i, contact := range contacts {
		contactsXML[i] = entity.ContactXML{
			Name: contact.Name,
			Numbers: []entity.ContactNumber{
				{
					Value: contact.Number},
			},
			Bandwidth:    0,
			Group:        1,
			Favorite:     0,
			Sticky:       0,
			Favoritetype: 0,
		}
	}
	contactList := entity.ContactList{Contacts: contactsXML}

	xmlFile, err := os.Create("contacts.xml")
	if err != nil {
		return fmt.Errorf("failed to create contacts.xml: %s", err)
	}
	enc := xml.NewEncoder(xmlFile)
	enc.Indent("  ", "    ")

	if err := enc.Encode(contactList); err != nil {
		return fmt.Errorf("failed to encode XML: %s", err)
	}

	return nil

}
