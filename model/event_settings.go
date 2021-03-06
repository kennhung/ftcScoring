package model

import (
	"time"
)

type EventSettings struct {
	Id                     int
	Name                   string
	Type                   string
	Region                 string
	Date                   time.Time
	DisplayBackgroundColor string
	DisplayOverlayMode     bool
}

const eventSettingsId = 0

func (database *Database) GetEventSettings() (*EventSettings, error) {
	eventSettings := new(EventSettings)
	err := database.eventSettingsMap.Get(eventSettings, eventSettingsId)
	if err != nil {
		// Database record doesn't exist yet; create it now.
		eventSettings.Name = "Untitled Event"
		eventSettings.Region = "Taiwan"
		eventSettings.DisplayBackgroundColor = "#00ff00"
		eventSettings.DisplayOverlayMode = true
		eventSettings.Date = time.Now()
		err = database.eventSettingsMap.Insert(eventSettings)
		if err != nil {
			return nil, err
		}
	}
	return eventSettings, nil
}

func (database *Database) SaveEventSettings(eventSettings *EventSettings) error {
	eventSettings.Id = eventSettingsId
	_, err := database.eventSettingsMap.Update(eventSettings)
	return err
}
