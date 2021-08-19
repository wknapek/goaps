package tournament

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

const (
	Win         = "win"
	Draw        = "draw"
	Loss        = "loss"
	TallyHeader = "Team                           | MP |  W |  D |  L |  P\n"
)

// Team holds the name and stats of a team.
type Team struct {
	Name           string
	MP, W, D, L, P int
}

// NewTeam returns reference to a new Team struct with name set to teamName.
func NewTeam(teamName string) *Team {
	return &Team{Name: teamName}
}

// Tally creates the scoreboard. Returns an error if creation is impeded.
func Tally(r io.Reader, w io.Writer) error {

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	teams, err := upadtateTeams(data)
	if err != nil {
		return err
	}

	result := TallyHeader
	for _, t := range teams {

		result += fmt.Sprintf("%-31s|  %d |  %d |  %d |  %d |  %d\n", t.Name, t.MP, t.W, t.D, t.L, t.P)
	}

	_, err = w.Write([]byte(result))

	return err
}

func upadtateTeams(data []byte) ([]*Team, error) {

	var teamsMap = make(map[string]*Team)

	rows := strings.Split(string(data), "\n")

	for _, row := range rows {
		if strings.HasPrefix(row, "#") || row == "" {
			continue
		}

		e := strings.Split(row, ";")

		if len(e) != 3 {
			return nil, errors.New("wrong number of fields")
		}

		team1Name, team2Name, result := e[0], e[1], e[2]

		if _, exist := teamsMap[team1Name]; !exist {
			teamsMap[team1Name] = NewTeam(team1Name)
		}
		if _, exist := teamsMap[team2Name]; !exist {
			teamsMap[team2Name] = NewTeam(team2Name)
		}

		t1, t2 := teamsMap[team1Name], teamsMap[team2Name]
		switch result {
		case Draw:
			t1.draws(t2)
		case Win:
			t1.wins(t2)
		case Loss:
			t2.wins(t1)
		default:
			return nil, errors.New("match result corrupted")
		}
		teamsMap[team1Name], teamsMap[team2Name] = t1, t2
	}

	var teams []*Team
	for _, t := range teamsMap {
		teams = append(teams, t)
	}

	sort.Slice(teams, func(i, j int) bool {
		if teams[i].P == teams[j].P {
			return teams[i].Name < teams[j].Name
		}
		return teams[i].P > teams[j].P
	})

	return teams, nil
}

// Methods
func (t *Team) draws(o *Team) {
	t.MP++
	o.MP++
	t.D++
	o.D++
	t.P++
	o.P++
}

func (t *Team) wins(o *Team) {
	t.MP++
	o.MP++
	t.W++
	t.P += 3
	o.L++
}
