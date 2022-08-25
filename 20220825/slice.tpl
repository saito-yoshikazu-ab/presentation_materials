{{- define "slice.tpl"}}
{{- $upperName := toUpperCase .Name}}
{{- $shortHandName := shortHand .Name}}
type {{$upperName}}s []*{{$upperName}}

func ({{$shortHandName}}s *{{$upperName}}s) ToMap() map[{{$upperName}}PK]*{{$upperName}} {
	m := make(map[{{$upperName}}PK]*{{$upperName}}, len(*{{$shortHandName}}s))
	for _, {{$shortHandName}} := range *{{$shortHandName}}s {
		m[{{$shortHandName}}.to{{$upperName}}PK()] = {{$shortHandName}}
	}
	return m
}
{{- end}}