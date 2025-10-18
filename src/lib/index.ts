// place files you want to import through the `$lib` alias in this folder.

export const API_URL = "https://openion.onrender.com"

export function shortenName(name: string): string {
  name = name.trim();
  let firstnameBegin = name[0];
  let lastspace = name.lastIndexOf(' ');
  let lastnameBegin;

  if (lastspace >= 1 && lastspace + 1 < name.length) lastnameBegin = name[lastspace + 1];

  return `${firstnameBegin} ${lastnameBegin}`;
}

