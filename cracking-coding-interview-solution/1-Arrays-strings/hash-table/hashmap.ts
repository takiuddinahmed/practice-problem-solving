class HashMap {
  private size: number;
  private data: string[][][] = [];

  constructor(size: number) {
    this.size = size;
  }

  hash(key: string) {
    return (
      key
        .split("")
        .reduce(
          (prevVal, currentVal, index) =>
            prevVal + currentVal.charCodeAt(0) * index,
          0
        ) % this.size
    );
  }

  insert(key: string, value: string) {
    const index = this.hash(key);
    if (!this.data[index]) {
      this.data[index] = [];
    }
    this.data[index].push([key, value]);
  }

  getValue(key: string) {
    const index = this.hash(key);
    if (!this.data[index] || !this.data[index].length) return null;
    return (
      this.data[index].filter((d) => d.length && d[0] === key)[0][1] || null
    );
  }

  getAll() {
    return this.data;
  }
}

const hashMap = new HashMap(10);
hashMap.insert("name", "Taki Uddin");
hashMap.insert("Address", "Bangladesh");
hashMap.insert("Phone", "+880123498714");

console.log(` Name -> ${hashMap.getValue("name")}`);
console.log(` Other -> ${hashMap.getValue("other")}`);
