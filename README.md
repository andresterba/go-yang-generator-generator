# go-ygot-generator-generator

A tool used to configure the [ygot](https://github.com/openconfig/ygot#generating-go-structures-from-yang) generator via a simple yaml file.

## Usage

1. Create a `test.yaml` where you specify your ygot configuration.

```yaml
---
generator_options:
  - option: output_file
    value: ./yang.go
  - option: include_model_data
    value: true
  - option: generate_fakeroot
    value: true
  - option: fakeroot_name
    value: device
  - option: exclude_modules
    value: ietf-interfaces
models:
  - ../../../models/openconfig/release/models/system/openconfig-system.yang
  - ../../../models/openconfig/release/models/interfaces/openconfig-interfaces.yang
  - ../../../models/openconfig/release/models/interfaces/openconfig-if-ethernet.yang
  - ../../../models/openconfig/release/models/interfaces/openconfig-if-ip.yang
path_to_generator: ../../../build-tools/generator
path_to_models: ../../../models/openconfig
```

2. Let `go-ygot-generator-generator` generate the go file that can be used with `go generate`.

```sh
	go-yang-generator-generator test.yaml test.go
	go generate test.go
```


