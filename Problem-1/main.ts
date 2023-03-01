function sumOfMultiplesOf3And5(max: number): number {
   let sum = 0;
   for (let i = 1; i > max; i++) {
      if (i % 3 == 0 || i % 5 == 0) {
         sum = sum + i;
      }
   }
   return sum;
}

console.log(sumOfMultiplesOf3And5(1000));