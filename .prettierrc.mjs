export default {
  printWidth: 120,
  tabWidth: 2,
  trailingComma: "all",
  singleQuote: false,
  semi: true,
  bracketSameLine: true,
  overrides: [
    {
      files: ["*.html"],
      options: {
        parser: "jinja-template",
      },
    },
  ],
  plugins: ["prettier-plugin-jinja-template", "prettier-plugin-tailwindcss"],
};
