CREATE TABLE "teams" (
  `id`          INTEGER PRIMARY KEY,
  `name`        VARCHAR(1000),
  `affiliation` VARCHAR(255),
  `city`        VARCHAR(255),
  `state`       VARCHAR(255),
  `country`     VARCHAR(255),
  `yellowcard`  bool
)