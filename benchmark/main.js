

let fib = (n) => {
    if (n < 0) {
      console.log("错误的输入");
    } else if (n == 0) {
      return 0;
    } else if (n == 1 || n == 2) {
      return 1;
    } else {
      return fib(n - 1) + fib(n - 2);
    }
}

 

start = Date.now() / 1000;
res = fib(35);
end = Date.now() / 1000;
engine = "nodejs"
console.log(`engine=${engine}, result=${res}, duration=${end-start}`)