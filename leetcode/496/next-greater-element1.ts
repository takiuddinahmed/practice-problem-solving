function nextGreaterElement(nums1: number[], nums2: number[]): number[] {
  const maxVal = Math.max(...nums2);
  const len = nums2.length;
  const result: number[] = [];
  for (const num of nums1) {
    if (num > maxVal) {
      result.push(-1);
      continue;
    }
    let matchedIndex = -1;
    for (let i = 0; i < len; i++) {
      if (num === nums2[i]) {
        matchedIndex = i;
        break;
      }
    }
    if (matchedIndex === -1) {
      result.push(-1);
      continue;
    }
    let found = false;
    for (let i = matchedIndex + 1; i < len; i++) {
      if (nums2[i] > num) {
        result.push(nums2[i]);
        found = true;
        break;
      }
    }
    if (!found) {
      result.push(-1);
    }
  }
  return result;
}
