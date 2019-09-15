package main

import "gopkg.in/yaml.v2"

func normalize(s string) (string, error) {
	t := map[string]interface{}{}
	if err := yaml.Unmarshal([]byte(s), &t); err != nil {
		return "", err
	}
	b, err := yaml.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(b), err
}
