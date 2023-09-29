const buffer = new Buffer.from("Tester", "utf-8");
buffer.write("Tester");

console.log(buffer);
console.log(buffer.toString());
console.log(buffer.toJSON());
