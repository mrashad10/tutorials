/// <reference path="./global.d.ts" />
//
// @ts-check
//
// The lines above enable type checking for this file. Various IDEs interpret
// the @ts-check and reference directives. Together, they give you helpful
// autocompletion when implementing this exercise. You don't need to understand
// them in order to use it.
//
// In your own projects, files, and code, you can play with @ts-check as well.

import { notify } from './notifier'
import { order } from './grocer'

/**
 * @return void
 */
export const onSuccess = () => notify({ message: 'SUCCESS' })

/**
 * @return void
 */
export const onError = () => notify({ message: 'ERROR' })

/**
 * @param {GrocerQuery} query
 * @param {FruitPickerSuccessCallback} onSuccessCallback
 * @param {FruitPickerErrorCallback} onErrorCallback
 * @return void
 */
export const orderFromGrocer = (query, onSuccessCallback, onErrorCallback) =>
  order(query, onSuccessCallback, onErrorCallback)

/**
 * @param {string} variety
 * @param {number} quantity
 * @return void
 */
export const postOrder = (variety, quantity) =>
  orderFromGrocer({ variety, quantity }, onSuccess, onError)
