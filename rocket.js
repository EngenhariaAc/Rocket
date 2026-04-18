class Rocket {
  constructor(name, fuel) {
    this.name = name;
    this.fuel = fuel;
  }

  launch() {
    console.log(`${this.name} is launching with ${this.fuel} fuel! 🚀`);
  }
}

module.exports = Rocket;
