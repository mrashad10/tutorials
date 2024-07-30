//
// This is only a SKELETON file for the 'Matrix' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export class Matrix {
  /**
 * Constructs a Matrix object.
 * @param {string} matrixString - The string representation of the matrix.
 */
  constructor(matrixString) {
    // Split the string by new line to get each line of the matrix
    const lines = matrixString.split('\n')

    // Map each line to an array of numbers
    this.matrix = lines.map(line => line.split(' ').map(Number))
  }

  /**
 * Returns the rows of the matrix.
 *
 * @returns {Array} The rows of the matrix.
 */
  get rows() {
    return this.matrix
  }

  /**
 * Get the columns of a matrix.
 * @returns {Array<Array<*>>} - The columns of the matrix.
 */
  get columns() {
    /*
     * Iterate over each row in the matrix. For each row, iterate over each
     * column in the row. Push the column value to the corresponding index in
     * the result array. If the column array doesn't exist in the result
     * array, initialize it as an empty array. Return the columns of the
     * matrix.
     * 
     * I prefer this solution than other solutions like:
     * return this.matrix[0].map((_, i) => this.matrix.map(row => row[i]))
     * because it can handle rows with different lengths.
     * */

    // Find the maximum length of any row in the matrix
    const maxLength = Math.max(...this.matrix.map(row => row.length))

    // Create an empty array to store the columns
    const result = []

    // Iterate through each row of the matrix
    this.matrix.forEach((row, index) => {
      // Iterate through each element in the row
      for (let i = 0; i < maxLength; i++) {
        // Create a new column if it doesn't exist
        result[i] = result[i] || []

        // Add the element to the column
        result[i][index] = row[i]
      }
    })

    // Return the columns of the matrix
    return result
  }
}
