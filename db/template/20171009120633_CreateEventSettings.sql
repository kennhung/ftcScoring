-- +goose Up
CREATE TABLE event_settings (
  id                     INTEGER PRIMARY KEY,
  name                   VARCHAR(255),
  type                   VARCHAR(16),
  region                 VARCHAR(16),
  date                   DATETIME,
  displaybackgroundcolor VARCHAR(16),
  displayoverlaymode     BOOLEAN
);

-- +goose Down
DROP TABLE event_settings;
