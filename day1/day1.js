const fs = require("fs");
const txt = fs.readFileSync("day1/input.txt");

let nos = txt.toString().split("\n"),ct=0;
nos = nos.map(a => parseInt(a));

//part 1
for(let i = 0; i < nos.length; i++){
    if(i == 0) continue;
    if(nos[i] > nos[i - 1]) ct++;
}
console.log(ct);

//part 2
ct = 0;
let window = nos[0] + nos[1] + nos[2];
for(j = 0; j < nos.length - 3; j++){
    if(window < window - nos[j] + nos[j+3]) 
        ct++;

    window = window - nos[j] + nos[j+3];
}
console.log(ct);