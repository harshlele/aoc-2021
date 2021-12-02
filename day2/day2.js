const fs = require("fs");
const txt = fs.readFileSync("day2/input.txt");

let lines = txt.toString().split("\r\n");

let pos = 0, depth = 0;
//part 1
for(let i = 0; i < lines.length; i++){
    let l = lines[i].split(" ");
    if(l[0] == "forward") pos += parseInt(l[1]);
    if(l[0] == "down") depth += parseInt(l[1]);
    if(l[0] == "up") depth -= parseInt(l[1]);
}
console.log(pos * depth);

//part 2
pos = 0;
depth = 0;
let aim = 0;

for(let i = 0; i < lines.length; i++){
    let l = lines[i].split(" ");
    
    if(l[0] == "forward") {
        pos += parseInt(l[1]);
        depth += aim * parseInt(l[1]);
    }
    if(l[0] == "down") aim += parseInt(l[1]);
    if(l[0] == "up") aim -= parseInt(l[1]);
}

console.log(pos * depth);

