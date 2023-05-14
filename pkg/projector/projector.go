package projector

import (
	"encoding/json"
	"os"
	"path"

	projector_config "github.com/zolwiastyl/submarine/pkg/config"
)

type Data struct {
	Projector map[string]map[string]string `json:"projector"`
}

type Projector struct {
	config *projector_config.Config
	data   *Data
}

func CreateProjector(config *projector_config.Config, data *Data) *Projector {
	return &Projector{config, data}
}

func (p *Projector) GetValue(key string) (string, bool) {
	curr := p.config.Pwd
	prev := ""
	out := ""
	found := false
	for prev != curr {

		if dir, ok := p.data.Projector[curr]; ok {
			if value, ok := dir[key]; ok {
				out = value
				found = true
				break
			}
		}
		prev = curr
		curr = path.Dir(curr)
	}

	return out, found
}

func (p *Projector) GetValueAll() map[string]string {
	out := map[string]string{}

	paths := []string{}

	curr := p.config.Pwd
	prev := ""

	for prev != curr {
		paths = append(paths, curr)

		prev = curr
		curr = path.Dir(curr)
	}

	for i := len(paths) - 1; i >= 0; i-- {
		if dir, ok := p.data.Projector[paths[i]]; ok {
			for k, v := range dir {
				out[k] = v
			}
		}
	}

	return out
}

func (p *Projector) SetValue(key, value string) {
	pwd := p.config.Pwd
	if _, ok := p.data.Projector[pwd]; !ok {
		p.data.Projector[pwd] = map[string]string{}
	}

	p.data.Projector[pwd][key] = value
}

func (p *Projector) RemoveValue(key string) {
	pwd := p.config.Pwd
	if dir, ok := p.data.Projector[pwd]; ok {
		delete(dir, key)
	}
}

func (p *Projector) Save() error {
	contents, err := json.Marshal(p.data)
	if err != nil {
		return err
	}
	return os.WriteFile(p.config.ConfigPath, contents, 0644)
}

func DefaultProjector(config *projector_config.Config) *Projector {
	return CreateProjector(config, &Data{
		Projector: map[string]map[string]string{},
	})
}

func NewProjector(config *projector_config.Config) *Projector {
	if _, err := os.Stat(config.ConfigPath); err != nil {
		contents, err := os.ReadFile(config.ConfigPath)
		if err != nil {
			return DefaultProjector(config)
		}
		var data Data
		err = json.Unmarshal(contents, &data)
		if err != nil {
			return DefaultProjector(config)
		}
		return &Projector{
			data:   &data,
			config: config,
		}
	}
	return DefaultProjector(config)
}
