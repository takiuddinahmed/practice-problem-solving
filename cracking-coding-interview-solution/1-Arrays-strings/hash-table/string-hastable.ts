class HashTable {
  private size: number;
  private data: string[][] = [];

  constructor(size: number) {
    this.size = size;
  }
  hash(value: string) {
    return (
      value.split("").reduce((previousValue, currentValue, currentIndex) => {
        return previousValue + currentValue.charCodeAt(0);
      }, 0) % this.size
    );
  }

  insert(value: string) {
    const index = this.hash(value);
    if (!this.data[index]) {
      this.data[index] = [];
    }
    this.data[index].push(value);
  }

  search(value: string) {
    const index = this.hash(value);
    if (!this.data[index] || !this.data[index].length) return null;
    return this.data[index].filter((d) => d === value)[0] || null;
  }

  getAllData() {
    return this.data.filter((d) => d).flat();
  }
}

const hashTable = new HashTable(10);
hashTable.insert("aaaa");
hashTable.insert("bbbb");
hashTable.insert("cccc");
hashTable.insert("dddd");
hashTable.insert("abbc");
hashTable.insert("ccda");
console.log(hashTable.search("cccc"));
console.log(hashTable.search("ssss"));

console.log(hashTable.getAllData());
