package model

import (
	"github.com/jmoiron/modl"
	"database/sql"
	"encoding/json"
	"path/filepath"
	"bitbucket.org/liamstask/goose/lib/goose"
)

const backupsDir = "db/backups"
const migrationsDir = "db/template"

type Database struct {
	Path             string
	db               *sql.DB
	eventSettingsMap *modl.DbMap
	matchMap         *modl.DbMap
	matchResultMap   *modl.DbMap
	rankingMap       *modl.DbMap
	teamMap          *modl.DbMap
	allianceTeamMap  *modl.DbMap
	lowerThirdMap    *modl.DbMap
	sponsorSlideMap  *modl.DbMap
}

func OpenDatabase(filename string) (*Database, error) {
	// Find and run the migrations using goose. This also auto-creates the DB.
	database := Database{Path: filename}
	migrationsPath := filepath.Join(".", migrationsDir)
	dbDriver := goose.DBDriver{"sqlite3", database.Path, "github.com/mattn/go-sqlite3", &goose.Sqlite3Dialect{}}
	dbConf := goose.DBConf{MigrationsDir: migrationsPath, Env: "prod", Driver: dbDriver}
	target, err := goose.GetMostRecentDBVersion(migrationsPath)
	if err != nil {
		return nil, err
	}
	err = goose.RunMigrations(&dbConf, migrationsPath, target)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", database.Path)
	if err != nil {
		return nil, err
	}
	database.db = db
	database.mapTables()

	return &database, nil
}

// Sets up table-object associations.
func (database *Database) mapTables() {
	dialect := new(modl.SqliteDialect)

	database.eventSettingsMap = modl.NewDbMap(database.db, dialect)
	database.eventSettingsMap.AddTableWithName(EventSettings{}, "event_settings").SetKeys(false, "Id")

	database.matchMap = modl.NewDbMap(database.db, dialect)
	database.matchMap.AddTableWithName(Match{}, "matches").SetKeys(true, "Id")

	database.matchResultMap = modl.NewDbMap(database.db, dialect)
	database.matchResultMap.AddTableWithName(MatchResultDb{}, "match_results").SetKeys(true, "Id")

	//TODO create DB
	database.rankingMap = modl.NewDbMap(database.db, dialect)
	database.rankingMap.AddTableWithName(RankingDb{}, "rankings").SetKeys(false, "TeamId")

	database.teamMap = modl.NewDbMap(database.db, dialect)
	database.teamMap.AddTableWithName(Team{}, "teams").SetKeys(false, "Id")
/*
	database.allianceTeamMap = modl.NewDbMap(database.db, dialect)
	database.allianceTeamMap.AddTableWithName(AllianceTeam{}, "alliance_teams").SetKeys(true, "Id")


	database.sponsorSlideMap = modl.NewDbMap(database.db, dialect)
	database.sponsorSlideMap.AddTableWithName(SponsorSlide{}, "sponsor_slides").SetKeys(true, "Id")
*/
}

func toJson(target *string, source interface{}) error {
	bytes, err := json.Marshal(source)
	if err != nil {
		return err
	}
	*target = string(bytes)
	return nil
}