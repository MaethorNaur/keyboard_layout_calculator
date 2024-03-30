package config

type Config struct {
	Languages map[string]Language
	Layouts   []Layout
	Scores    map[string][][]float64
}

type Language struct {
	Corpuses []string
	Letters  string
}

type Layout struct {
	Name      string
	Languages map[string]LayoutLanguage
}

type LayoutLanguage struct {
	Weight float64
	Scores string
}
