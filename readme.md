# coparse

> Parse CLI flags from a file


[![CI](https://github.com/barelyhuman/coparse/actions/workflows/ci.yml/badge.svg)](https://github.com/barelyhuman/coparse/actions/workflows/ci.yml)


> **Warning**: Still in ideation stage

> **Note**: This is being written to help with
> [commitlog](https://github.com/barelyhuman/commitlog)'s specific requirements
> and the style of configuration I wish to support for most of my apps. If you
> wish for anything else to be added please feel free to raise an issue

## Ideation

### Things to do

- Parse the passed file into flags
- Merge other flags with these flags
- Be able to define the hierarchy for overriding (default to, local file first,
  global later)

### Syntax

- plaintext file
- newline after each flag
- support `=`,`:`,`<space>` as the value and argument delimeter

```
--flag=false
--flag false
--flag:false
```

### Possible API

- `ParseConfig(pathToConfig string)`
- `MergeConfig(lowestPriority....highestPriority)`, so `lowestPriority` flags
  come from let's say a config file in the project and highestPriority is a flag
  override sent to the CLI directly.
- `FlagsToConfig`, something to convert flag variables into a config object to
  be used
