function containsDuplicate(nums: number[]): boolean {
  const map = new Map();
  for (const num of nums) {
    if (map.has(num)) {
      return true;
    }
    map.set(num, 1);
  }
  return false;
}

console.log(containsDuplicate([1, 2, 3, 1]));
console.log(containsDuplicate([1, 2, 3, 4]));
console.log(containsDuplicate([1, 1, 1, 3, 3, 4, 3, 2, 4, 2]));
