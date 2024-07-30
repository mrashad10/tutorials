/// <reference path="./global.d.ts" />
// @ts-check

/**
 * Get the first card in the given deck
 *
 * @param {Card[]} deck
 *
 * @returns {Card} the first card in the deck
 */
export const getFirstCard = ([firstCard]) => firstCard

/**
 * Get the second card in the given deck
 *
 * @param {Card[]} deck
 *
 * @returns {Card} the second card in the deck
 */
export const getSecondCard = ([, secondCard]) => secondCard

/**
 * Switch the position of the first two cards in the given deck
 *
 * @param {Card[]} deck
 *
 * @returns {Card[]} new deck with reordered cards
 */
export const swapTopTwoCards = ([firstCard, secondCard, ...everythingElse]) => [
  secondCard,
  firstCard,
  ...everythingElse
]

/**
 * Put the top card of the given deck into a separate discard pile
 *
 * @param {Card[]} deck
 *
 * @returns {[Card, Card[]]} the top card of the given
 * deck and a new deck containing all the other cards
 */
export const discardTopCard = ([firstCard, ...everythingElse]) => [
  firstCard,
  everythingElse
]

/** @type {Card[]} **/
const FACE_CARDS = ['jack', 'queen', 'king']

/**
 * Insert face cards into the given deck
 *
 * @param {Card[]} deck
 *
 * @returns {Card[]} new deck where the second,
 * third, and fourth cards are the face cards
 */
export const insertFaceCards = ([firstCard, ...everythingElse]) => [
  firstCard,
  'jack',
  'queen',
  'king',
  ...everythingElse
]
