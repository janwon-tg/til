const message: string = "Hello, world!";
const message2: string = "Hello, world2!";

console.log(message);
console.log(message2);
console.log(range(5, 10.0));
helloNtimes(5);

function helloNtimes(n: number): void {
  for (let i = 0; i < n; i++) {
    console.log("Hello World!!");
  }
}

function range(min: number, max: number): number[] {
  const result = [];
  for (let i = min; i <= max; i++) {
    result.push(i);
  }
  return result;
}


