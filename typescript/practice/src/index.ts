const message: string = "Hello, world!";
const message2: string = "Hello, world2!";

type Human = {
  height: number;
  weight: number;
};

interface Human2 {
  height: number;
  weight: number;
}

console.log(message);
console.log("Hello");
console.log("Hello");
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

/*
   src/index.ts:33:28 - error TS2339: Property 'h' does not exist on type 'Human'.
   33 const calcBMI3 = function({ h, w }: Human): number {
   src/index.ts:33:31 - error TS2339: Property 'w' does not exist on type 'Human'.
   33 const calcBMI = function({ h, w }: Human): number {
*/
//const calcBMI3 = function({ h, w }: Human): number {
//	return weight / height ** 2;
//};
const calcBMI = function ({ height, weight }: Human): number {
  return weight / height ** 2;
};

const calcBMI2 = function (h: Human): number {
  return h.weight / h.height ** 2;
};

const calcBMI4 = ({ height, weight }: Human): number => {
  return weight / height ** 2;
};

const calcBMI5 = ({ height, weight }: Human): number => weight / height ** 2;

const taro: Human = { height: 1.8, weight: 80 };
const jiro: Human2 = { height: 1.3, weight: 36 };

console.log(calcBMI(taro));
console.log(calcBMI(jiro));
console.log(calcBMI({ height: 1.7, weight: 77 }));

console.log(calcBMI2(taro));
console.log(calcBMI2(jiro));
console.log(calcBMI2({ height: 1.7, weight: 77 }));

console.log(calcBMI4(taro));
console.log(calcBMI4(jiro));
console.log(calcBMI4({ height: 1.7, weight: 77 }));

console.log(calcBMI5(taro));
console.log(calcBMI5(jiro));
console.log(calcBMI5({ height: 1.7, weight: 77 }));
