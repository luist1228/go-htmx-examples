package util

import "context"

type Theme struct {
	Name  string
	Value string
}

var Themes []Theme = []Theme{
	{
		Name:  "light",
		Value: "light",
	},
	{
		Name:  "dark",
		Value: "dark",
	},
	{
		Name:  "cupcake",
		Value: "cupcake",
	},
	{
		Name:  "bumblebee",
		Value: "bumblebee",
	},
	{
		Name:  "emerald",
		Value: "emerald",
	},
	{
		Name:  "corporate",
		Value: "corporate",
	},
	{
		Name:  "synthwave",
		Value: "synthwave",
	},
	{
		Name:  "retro",
		Value: "retro",
	},
	{
		Name:  "cyberpunk",
		Value: "cyberpunk",
	},
	{
		Name:  "valentine",
		Value: "valentine",
	},
	{
		Name:  "halloween",
		Value: "halloween",
	},
	{
		Name:  "garden",
		Value: "garden",
	},
	{
		Name:  "forest",
		Value: "forest",
	},
	{
		Name:  "aqua",
		Value: "aqua",
	},
	{
		Name:  "lofi",
		Value: "lofi",
	},
	{
		Name:  "pastel",
		Value: "pastel",
	},
	{
		Name:  "fantasy",
		Value: "fantasy",
	},
	{
		Name:  "wireframe",
		Value: "wireframe",
	},
	{
		Name:  "black",
		Value: "black",
	},
	{
		Name:  "luxury",
		Value: "luxury",
	},
	{
		Name:  "dracula",
		Value: "dracula",
	},
	{
		Name:  "cmyk",
		Value: "cmyk",
	},
	{
		Name:  "autumn",
		Value: "autumn",
	},
	{
		Name:  "business",
		Value: "business",
	},
	{
		Name:  "acid",
		Value: "acid",
	},
	{
		Name:  "lemonade",
		Value: "lemonade",
	},
	{
		Name:  "night",
		Value: "night",
	},
	{
		Name:  "coffee",
		Value: "coffee",
	},
	{
		Name:  "winter",
		Value: "winter",
	},
	{
		Name:  "dim",
		Value: "dim",
	},
	{
		Name:  "nord",
		Value: "nord",
	},
	{
		Name:  "sunset",
		Value: "sunset",
	},
}


func TemplIsHtmx(ctx context.Context) bool {
	return ctx.Value("isHtmx").(bool)
}