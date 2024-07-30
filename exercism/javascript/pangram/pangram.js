//
// This is only a SKELETON file for the 'Pangram' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

// Check if the given sentence is a pangram
export const newIsPangram = sentence => {
  // Create an array of booleans to keep track of the presence of each letter
  const letters = new Array(26).fill(false)

  // Convert the sentence to lowercase
  const lowercaseSentence = sentence.toLowerCase()

  // Iterate through each character in the sentence
  for (let i = 0; i < lowercaseSentence.length; i++) {
    // Get the character code of the current character
    const charCode = lowercaseSentence.charCodeAt(i)

    // Check if the character is a lowercase letter (a-z)
    if (charCode >= 97 && charCode <= 122) {
      // Calculate the index of the letter in the letters array
      const index = charCode - 97

      // Set the corresponding letter flag to true
      letters[index] = true
    }
  }

  // Check if all letters are present (i.e., all flags are true)
  return letters.every(letter => letter)
}

export const oldIsPangram = sentence =>
  new Set([...sentence.toLowerCase()].filter(x => /[a-z]/.test(x))).size === 26
