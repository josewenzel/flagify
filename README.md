# Flagify

Self-hosted feature flags service written in Go

## How to run locally

## Data schemas

### Project

A project is an aggregation of feature flags, it serves as a way to group feature flags that may be used by the same 
service or module.

```json
{
  "id": "string",
  "name": "string",
  "flags": []
}
```

### Flag

A flag is a toggle for a feature that may be on or off, meaning that you can disable pieces of your code with a more 
meaningful attribute. They also support randomized canary releases, which means that you can turn that feature on, only
to a percentage group of users (currently there is no way to track which users were given true or false, so be mindful
if using for AB testing).

```json
{
  "id": "string",
  "name": "string",
  "value": "boolean",
  "canarySetting": "float"
}
```