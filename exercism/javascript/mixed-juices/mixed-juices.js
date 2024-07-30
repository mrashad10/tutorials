// @ts-check
//
// The line above enables type checking for this file. Various IDEs interpret
// the @ts-check directive. It will give you helpful autocompletion when
// implementing this exercise.

/**
 * Determines how long it takes to prepare a certain juice.
 *
 * @param {string} name
 * @returns {number} time in minutes
 */
export const timeToMixJuice = name => {
  switch (name) {
    case 'Pure Strawberry Joy':
      return 0.5

    case 'Energizer':
    case 'Green Garden':
      return 1.5

    case 'Tropical Island':
      return 3

    case 'All or Nothing':
      return 5

    default:
      return 2.5
  }
}

/**
 * Calculates the number of limes that need to be cut
 * to reach a certain supply.
 *
 * @param {number} wedgesNeeded
 * @param {string[]} limes
 * @returns {number} number of limes cut
 */
export const limesToCut = (wedgesNeeded, limes) => {
  if (wedgesNeeded === 0 || limes.length === 0) return 0
  const counts = { small: 6, medium: 8, large: 10 }
  let index = 0
  let count = 0

  while (index < limes.length && wedgesNeeded > count)
    count += counts[limes[index++]]

  return index
}

/**
 * Determines which juices still need to be prepared after the end of the shift.
 *
 * @param {number} timeLeft
 * @param {string[]} orders
 * @returns {string[]} remaining orders after the time is up
 */
export const remainingOrders = (timeLeft, orders) =>
  orders
    .map(order => {
      if (timeLeft > 0) {
        timeLeft -= timeToMixJuice(order)
        return ''
      } else {
        return order
      }
    })
    .filter(order => order !== '')
