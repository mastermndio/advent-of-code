const input = `paste input here`

function inputParse(inputString) {
  splitInput = inputString.split("\n")
  let inputInts = []
  for (let i = 0; i < splitInput.length; i++) {
    inputInts.push(Number(splitInput[i]))
  }

  return inputInts
}
