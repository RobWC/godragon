package godragon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

// StaticChampions returns a map of all Champions
func StaticChampions(version string) (cr map[string]Champion, err error) {
	path := fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/data/en_US/champion.json", version)
	resp, err := http.Get(path)
	if err != nil {
		return cr, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return cr, err
	}
	cd := ChampionData{}
	json.Unmarshal(body, &cd)
	return cd.Champions, err
}

// StaticChampiom returns a specific champion
func StaticChampion(name string, version string) (c Champion, err error) {
	path := fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/data/en_US/champion/%s.json", version, name)
	resp, err := http.Get(path)
	if err != nil {
		return c, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c, err
	}
	cd := ChampionData{}
	json.Unmarshal(body, &cd)
	return cd.Champions[name], err
}

func add(a, b int) int {
	return a + b
}

func mult(a int, b float32) float32 {
	return float32(a) * b
}

func ascalc(aso float32) string {
	s := math.Pow(10, float64(3))
	v := float64(0.625 / (math.Floor((1-float64(aso))*s) / s))
	nv := strconv.FormatFloat(v, 'f', -1, 32)
	return nv[:5]
}

// ChampionCooldowns output the champion cooldowns to a formatted string
func ChampionCooldowns(c Champion) error {

	fmap := template.FuncMap{"add": add, "mult": mult, "ascalc": ascalc}

	output :=
		`{{ .Name }} --- {{ .Title }}
Type: {{$tlen := len .Tags}}{{ range $i, $e := .Tags}}{{$e}}{{ $v := add $i 1}}{{if ne $v $tlen}}, {{end}}{{end}}
Skins: {{ len .Skins }}

HP:    {{ .Stats.HP }}(+{{.Stats.HPPerLevel}})   HP Regen:   {{.Stats.HPRegen}}(+{{.Stats.HPRegenPerLevel}}) 
Mana:  {{ .Stats.MP }}(+{{.Stats.MPPerLevel}})   Mana Regen: {{ .Stats.MPRegen}}(+{{.Stats.MPRegenPerLevel}})
Armor: {{ .Stats.Armor}}(+{{.Stats.ArmorPerLevel}})     MR:         {{.Stats.SpellBlock}}(+{{.Stats.SpellBlockPerLevel}})
AD:    {{ .Stats.AttackDamage}}(+{{.Stats.AttackDamagePerLevel}})  AS:         {{ ascalc .Stats.AttackSpeedOffset}}(+{{.Stats.AttackSpeedPerLevel}})
Crit:  {{ .Stats.Crit}}(+{{.Stats.CritPerLevel}})         Range:         {{ .Stats.AttackRange}}
MS:    {{ .Stats.MoveSpeed }}                       

Spells:
------------------------------------------
{{printf "%-22s" .Passive.Name}}    Passive
{{ range $i, $v := .Spells }}{{ printf "%-22s" $v.Name }}    {{ $tlen := len .Cooldown }}{{ range $i, $e := .Cooldown}}{{$v := add $i 1}}{{ $e }}{{ if ne $v $tlen}}/{{end }}{{ end }}
{{ end  }}
`

	t := template.Must(template.New("champ").Funcs(fmap).Parse(output))

	err := t.Execute(os.Stdout, c)
	if err != nil {
		return err
	}

	return nil
}
