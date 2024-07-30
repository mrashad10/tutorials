//
// This is only a SKELETON file for the 'Space Age' exercise. It's been
// provided as a convenience to get you started writing code faster.
//

/**
 * Calculate the age on a specific planet based on the given seconds.
 *
 * @param {string} planet - The name of the planet.
 * @param {number} seconds - The number of seconds.
 * @return {number} The age on the specified planet.
 */
export const age = (planet, seconds) => {
  // Number of seconds in an Earth year
  const earthYearInSeconds = 31557600

  // Conversion rates for different planets
  const conversionRates = {
    mercury: earthYearInSeconds * 0.2408467,
    venus: earthYearInSeconds * 0.61519726,
    earth: earthYearInSeconds,
    mars: earthYearInSeconds * 1.8808158,
    jupiter: earthYearInSeconds * 11.862615,
    saturn: earthYearInSeconds * 29.447498,
    uranus: earthYearInSeconds * 84.016846,
    neptune: earthYearInSeconds * 164.79132
  }

  // Calculate the age on the specified planet
  const ageOnPlanet = seconds / conversionRates[planet]

  // Return the age on the specified planet, rounded to 2 decimal places
  return parseFloat(ageOnPlanet.toFixed(2))
}
