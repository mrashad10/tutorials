// @ts-check

// Use a precompiled regular expression: Instead of creating the regex object
// inside the functions on every function call, This will save the overhead of
// creating the regex object repeatedly.

// Regex for matching a command
const commandRegex = /^chatbot/i

// Regex for matching an emoji
// Using the RegExp constructor instead of the RegExp literal to satisfy the
// requirements of the exercise.
const emojiRegex = new RegExp("emoji[0-9]+", "g")

// Regex for matching a phone number
const phoneNumberRegex = /\(\+\d{2}\) \d{3}-\d{3}-\d{3}/

// Regex for matching a websites
const websiteRegex = /([a-z0-9.-]+\.[a-z]{2,})/gi

// Regex for extracting the frist and last names from a full name
const fullNameRegex = /([^,]+),\s*(.+)/

/**
 * Given a certain command, help the chatbot recognize whether the command is valid or not.
 *
 * @param {string} command
 * @returns {boolean} whether or not is the command valid
 */

export const isValidCommand = command => commandRegex.test(command)

/**
 * Given a certain message, help the chatbot get rid of all the emoji's encryption through the message.
 *
 * @param {string} message
 * @returns {string} The message without the emojis encryption
 */
export const removeEmoji = message => message.replace(emojiRegex, "")

/**
 * Given a certain phone number, help the chatbot recognize whether it is in the correct format.
 *
 * @param {string} number
 * @returns {string} the Chatbot response to the phone Validation
 */
export const checkPhoneNumber = number =>
  phoneNumberRegex.test(number)
    ? "Thanks! You can now download me to your phone."
    : `Oops, it seems like I can't reach out to ${number}`

/**
 * Given a certain response from the user, help the chatbot get only the URL.
 *
 * @param {string} userInput
 * @returns {string[] | null} all the possible URL's that the user may have answered
 */
export const getURL = userInput => userInput.match(websiteRegex)

/**
 * Greet the user using the full name data from the profile.
 *
 * @param {string} fullName
 * @returns {string} Greeting from the chatbot
 */
export const niceToMeetYou = fullName =>
  fullName.replace(fullNameRegex, "Nice to meet you, $2 $1")
