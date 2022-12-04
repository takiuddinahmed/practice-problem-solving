/*
Problem statement: Implement an algorithm to determine if a string has all unique characters.
*/

function isUniqueCharacters(value: string) {
  const valObject: { [key: string]: number } = {};
  for (const char of value) {
    if (valObject[char]) {
      return false;
    }
    valObject[char] = 1;
  }
  return true;
}

console.log(
  `abcdef -> ${isUniqueCharacters("abcdef") ? "unique" : "not-unique"}`
);
console.log(
  `abcdaf -> ${isUniqueCharacters("abcdaf") ? "unique" : "not-unique"}`
);
console.log(
  `Hello -> ${isUniqueCharacters("Hello") ? "unique" : "not-unique"}`
);
console.log(
  `bahdc -> ${isUniqueCharacters("bahdc") ? "unique" : "not-unique"}`
);
