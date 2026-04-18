class Engine:
    def __init__(self, thrust, fuel_type):
        self.thrust = thrust
        self.fuel_type = fuel_type

    def ignite(self):
        print(f"Engine ignited with {self.thrust} kN thrust using {self.fuel_type}.")
