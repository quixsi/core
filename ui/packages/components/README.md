# @quixsi/components

A collection of UI components generated from [shadcn-vue](https://www.shadcn-vue.com/).

## Installation

> [!IMPORTANT]
> This package is a private package and not put on a registry.
> Run the following command from the workspace root and make sure to use `*` as version.

```sh
npm i -w <workspace> ./packages/components
```

or manually edit the manifest and run installation

```json
{
  "dependencies": {
    "@quixsi/components": "*"
  }
}
```

## How to use

After installing the package in the project each component can be imported individually or all at once.

```ts
import { Button } from "@quixsi/components/button"
import { Input } from "@quixsi/components/input"
import { Button, Input } from "@quixsi/components"
```

It is recommended to register components globally when used in a lot of different views or features.

```ts
import { createApp } from 'vue'
import { Button } from "@quixsi/components/button"

const app = createApp(App)

app.component('Button', Button)

// ...
```

## Contributing

Use the [CLI](https://www.shadcn-vue.com/docs/cli.html) to add more components.

> [!IMPORTANT]
> From the root of the repository or UI project, make sure to use the `--cwd` (`-c`).
> Example: `npx shadcn-vue@latest add button -c ./packages/components`
