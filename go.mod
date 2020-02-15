module github.com/asobti/kube-monkey

go 1.13

require (
	github.com/bouk/monkey v0.0.0-00010101000000-000000000000
	github.com/davecgh/go-spew v1.1.1
	github.com/fsnotify/fsnotify v1.4.2
	github.com/ghodss/yaml v0.0.0-20150909031657-73d445a93680
	github.com/go-kit/kit v0.9.0
	github.com/go-logfmt/logfmt v0.5.0 // indirect
	github.com/gogo/protobuf v0.0.0-20170330071051-c0656edd0d9e
	github.com/golang/glog v0.0.0-20141105023935-44145f04b68c
	github.com/golang/protobuf v0.0.0-20171021043952-1643683e1b54
	github.com/google/gofuzz v0.0.0-20161122191042-44d81051d367
	github.com/googleapis/gnostic v0.0.0-20170729233727-0c5108395e2d
	github.com/hashicorp/hcl v0.0.0-20180404174102-ef8a98b0bbce
	github.com/json-iterator/go v0.0.0-20171212105241-13f86432b882
	github.com/magiconair/properties v1.8.0
	github.com/mitchellh/mapstructure v0.0.0-20180511142126-bb74f1db0675
	github.com/pelletier/go-toml v1.2.0
	github.com/pkg/errors v0.8.0
	github.com/pmezard/go-difflib v1.0.0
	github.com/spf13/afero v1.1.0
	github.com/spf13/cast v1.2.0
	github.com/spf13/jwalterweatherman v0.0.0-20180109140146-7c0cea34c8ec
	github.com/spf13/pflag v1.0.1-0.20171106142849-4c012f6dcd95
	github.com/spf13/viper v1.0.2
	github.com/stretchr/objx v0.1.2-0.20180531200725-0ab728f62c7f
	github.com/stretchr/testify v1.2.2-0.20180319223459-c679ae2cc0cb
	golang.org/x/crypto v0.0.0-20170825220121-81e90905daef
	golang.org/x/net v0.0.0-20170809000501-1c05540f6879
	golang.org/x/sys v0.0.0-20171031081856-95c657629925
	golang.org/x/text v0.0.0-20170810154203-b19bf474d317
	golang.org/x/time v0.0.0-20161028155119-f51c12702a4d
	gopkg.in/inf.v0 v0.9.0
	gopkg.in/yaml.v2 v2.0.0-20170721113624-670d4cfef054
	k8s.io/api v0.0.0-20180308224125-73d903622b73
	k8s.io/apimachinery v0.0.0-20180228050457-302974c03f7e
	k8s.io/client-go v7.0.0+incompatible
)

replace github.com/bouk/monkey => bou.ke/monkey v1.0.2
