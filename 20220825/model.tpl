package modelrepository

//go:generate mockgen -destinition=mock_$GOFILE -package=$GOPACKAGE

import (
    "time"

    "github.com/jmoiron/sqlx"
)
{{$base := .}}
{{$upperName := toUpperCase .Name}}
type {{$upperName}} struct {
    {{- range $key, $element := .Fields}}
    {{$element.Name}} {{$element.Type}}
    {{- end}}
}

type {{$upperName}}PK struct {
    {{- range $index, $element := .Fields}}
    {{- if $element.IsPK}}
    {{$element.Name}} {{$element.Type}}
    {{- end}}
    {{- end}}
}
{{$shortHandName := shortHand .Name}}
func ({{$shortHandName}} *{{$upperName}}) to{{$upperName}}PK() {{$upperName}}PK {
	return {{$upperName}}PK{
        {{- range $index, $element := .Fields}}
        {{- if $element.IsPK}}
        {{$element.Name}}: {{$shortHandName}}.{{$element.Name}},
        {{- end}}
        {{- end}}
    }
}

{{template "slice.tpl" $base}}

type {{$upperName}}ModelRepository struct {
    client *sqlx.DB
}

func (r *{{$upperName}}ModelRepository) Get({{$shortHandName}}k {{$upperName}}PK) (*{{$upperName}}, error) {
	model := new({{$upperName}})
	if err := r.client.Select(&model, `SELECT * FROM {{.Name}} WHERE
    {{- $fields := .Fields}}
    {{- range $index, $element := $fields}}
    {{- if $element.IsPK}}
    {{$element.Name}}=?
    {{- if ne $index (pkNum $fields | subOne)}}
    AND
    {{- end}}
    {{- end}}
    {{- end}}
    `,
    {{- range $index, $element := .Fields}}
    {{- if $element.IsPK}}
    {{$shortHandName}}k.{{$element.Name}},
    {{- end}}
    {{- end}}
    ); err != nil {
		return nil, err
	}
	return model, nil
}

{{- range $index, $element := .Fields}}
{{- if $element.IsSK}}
{{$lowerName := toLowerCase $element.Name}}
func (r *{{$upperName}}ModelRepository) FindBy{{$element.Name}}({{$lowerName}} string) ({{$upperName}}s, error) {
	var models {{$upperName}}s
	if err := r.client.Select(&models, "SELECT * FROM {{$base.Name}} WHERE {{$lowerName}}=?", {{$lowerName}}); err != nil {
		return nil, err
	}
	return models, nil
}
{{- end}}
{{- end}}
