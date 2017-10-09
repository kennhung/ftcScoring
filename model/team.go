package model

type Team struct {
	Id              int
	Name            string
	Affiliation     string
	City            string
	State           string
	Country         string
	YellowCard      bool
}

func (database *Database) CreateTeam(team *Team) error {
	return database.teamMap.Insert(team)
}

func (database *Database) GetTeamById(id int) (*Team, error) {
	team := new(Team)
	err := database.teamMap.Get(team, id)
	if err != nil && err.Error() == "sql: no rows in result set" {
		team = nil
		err = nil
	}
	return team, err
}

func (database *Database) SaveTeam(team *Team) error {
	_, err := database.teamMap.Update(team)
	return err
}

func (database *Database) DeleteTeam(team *Team) error {
	_, err := database.teamMap.Delete(team)
	return err
}

func (database *Database) TruncateTeams() error {
	return database.teamMap.TruncateTables()
}

func (database *Database) GetAllTeams() ([]Team, error) {
	var teams []Team
	err := database.teamMap.Select(&teams, "SELECT * FROM teams ORDER BY id")
	return teams, err
}
