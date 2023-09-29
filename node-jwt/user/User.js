/**
 * @module User
 * @description User model for the application
 */

const mongoose = require("mongoose");
const bcrypt = require("bcrypt");

/**
 * @typedef {Object} UserSchema
 * @property {String} name - Name of the user
 * @property {String} email - Email of the user
 * @property {String} password - Hashed password of the user
 */

const UserSchema = new mongoose.Schema({
  name: {
    type: String,
    required: true,
    trim: true,
  },
  email: {
    type: String,
    required: true,
    unique: true,
    trim: true,
  },
  password: {
    type: String,
    required: true,
    minlength: 8,
  },
});

/**
 * @description Middleware to hash the password before saving the user
 */
UserSchema.pre("save", async function (next) {
  if (this.isModified("password")) {
    this.password = await bcrypt.hash(this.password, 10);
  }
  next();
});

/**
 * @description Method to compare the password entered by the user with the hashed password stored in the database
 * @param {String} enteredPassword - Password entered by the user
 * @returns {Promise<boolean>} - Promise that resolves to true if the password is correct, false otherwise
 */
UserSchema.methods.comparePassword = function (enteredPassword) {
  return bcrypt.compare(enteredPassword, this.password);
};

const User = mongoose.model("User", UserSchema);

module.exports = User;
