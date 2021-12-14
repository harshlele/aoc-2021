//some rough work for the actual solution, which is in go lol


let str = "NNCB"

let o = {
    "CH": "B",
    "HH": "N",
    "CB": "H",
    "NH": "C",
    "HB": "C",
    "HC": "B",
    "HN": "C",
    "NN": "C",
    "BH": "H",
    "NC": "B",
    "NB": "B",
    "BN": "B",
    "BB": "N",
    "BC": "B",
    "CC": "N",
    "CN": "C"
}

//let comb = { NC: 1, CN: 1, NB: 1, BC: 1, CH: 1, HB: 1 }
let comb = { NB: 2, BC: 2, CC: 1, CN: 1, BB: 2, CB: 2, BH: 1, HC: 1 }

let newO = {

}

Object.keys(comb).forEach((k,i) => {
    if (o[k]){
        let p = o[k]
        let [k1,k2] = k.split("")

        if(newO[`${k1}${p}`]){
            newO[`${k1}${p}`] += comb[k]
        }
        else {
            newO[`${k1}${p}`] = comb[k]
        }

        if(newO[`${p}${k2}`]){
            newO[`${p}${k2}`] += comb[k]
        }
        else {
            newO[`${p}${k2}`] = comb[k]
        }
    }
});


console.log(newO)