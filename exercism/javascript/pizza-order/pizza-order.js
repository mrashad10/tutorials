/// <reference path="./global.d.ts" />
//
// @ts-check

/**
 * Determine the prize of the pizza given the pizza and optional extras
 *
 * @param {Pizza} pizza name of the pizza to be made
 * @param {Extra[]} extras list of extras
 *
 * @returns {number} the price of the pizza
 */
const pizzaPrice = (pizza, ...extras) => {
  const prices = {
    Margherita: 7,
    Caprese: 9,
    Formaggio: 10,
  };

  const extraPrices = {
    ExtraSauce: 1,
    ExtraToppings: 2,
  };

  let price = prices[pizza] || 0;

  extras.forEach(extra => {
    price += extraPrices[extra] || 0;
  });

  return price;
}

/**
 * Calculate the prize of the total order, given individual orders
 *
 * @param {PizzaOrder[]} pizzaOrders a list of pizza orders
 * @returns {number} the price of the total order
 */
const orderPrice = (pizzaOrders) => pizzaOrders.reduce((total, pizzaOrder) => {
  return total + pizzaPrice(pizzaOrder.pizza, ...pizzaOrder.extras);
}, 0);



module.exports = {
  pizzaPrice, orderPrice
}