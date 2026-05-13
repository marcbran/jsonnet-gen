local p = import 'pkg/main.libsonnet';

p.pkg({
  source: 'https://github.com/marcbran/jsonnet-libs',
  repo: 'https://github.com/marcbran/jsonnet.git',
  branch: 'promql',
  path: 'promql',
  target: 'p',
}, |||
  Jsonnet library that implements a DSL for PromQL.
|||)
