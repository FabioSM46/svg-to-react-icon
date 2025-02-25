# svg-to-tsx

A CLI tool to convert SVG files into TypeScript React functional components.

## Features

- Converts SVG files into TSX components.
- Supports both filled and outlined versions.
- Converts dashed attributes to camelCase.
- Applies props like `size`, `color`, and `strokeWidth` to SVG elements.
- Generates an `index.ts` file to export all components.
- Processes two input folders containing matching SVG files.
- Ensures filenames start with `_` if they begin with a number.

## Installation

Clone the repository and build the binary:

```sh
go build -o svg-to-tsx .
```

## Usage

Run the CLI with two folders as input:

```sh
./svg-to-tsx ./filled-icons ./outlined-icons
```

### Example

If you have:

```
filled-icons/
  home.svg
  search.svg

outlined-icons/
  home.svg
  search.svg
```

The tool will generate:

```
dist/
  Home.tsx
  Search.tsx
  index.ts
```

## Output Format

Each generated component follows this structure:

```tsx
import React from "react";
import { SVGProps } from "react";

const Home = (props: SVGProps<SVGSVGElement>) => (
  <svg {...props} viewBox='0 0 24 24'>
    <path d='M...' />
  </svg>
);

export default Home;
```

## Contributing

Feel free to open issues or submit pull requests!

## License

MIT License
