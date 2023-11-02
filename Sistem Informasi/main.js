const readline = require('readline');

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

let x;
let i = 0;
let a = 7;
let b = 9;
let c = 100;
let f = c % 3;

rl.question("Masukkan nilai x: ", (inputX) => {
    x = parseInt(inputX);
    rl.close();

    if (!Number.isInteger(x)) {
        console.log("Nilai x harus integer.");
        return;
    }

    for (i = 0; i < 100; i += 10) {
        x -= 20;
        b -= 2;

        if (b > 0) {
            b++;
            a++;
            f += a;
        } else {
            a++;
            f += a;
        }

        for (; f % 4 === 0;) {
            b++;
            a++;
            f += a;
        }
    }

    b *= 2;

    console.log(`Value a: ${a}\nValue b: ${b}\nValue c: ${c}\nValue f: ${f}\nValue x: ${x}`);
});
