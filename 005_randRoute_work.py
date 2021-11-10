import random


class UniversalTrip(object):

    def __init__(self):
        self.spaceline = ["Space Adventures", "SpaceX", "Virgin Galactic"]
        self.tripType = ["Round-trip", "One-way"]

    def tripTable(self):
        print("%-20s      %-10s    %-10s      %-10s" %("Spaceline", "Days", "Triptype", "Price"))
        print("=" * 70)
        for i in range(1, 10):
            trip = self.tripType[random.randint(0, 1)]
            days = self.tripDays(trip)
            price = self.tripPrice(trip)
            print("%-20s      %-10d    %-10s      $ %-10d" % (self.spaceline[random.randint(0, 2)], days, trip, price))
            
    @staticmethod
    def tripDays(trip):
        speed = random.randint(16, 30)
        distance = 62100000
        days = int(distance/speed) + 1
        return days
        if trip == "One-way":
            return days
        else:
            return days * 2

    @staticmethod
    def tripPrice(trip):

        price = random.randint(3600, 5000)  # 3600-5000
        if trip == "One-way":
            return price
        else:
            return price * 2


if __name__ == "__main__":

    table = UniversalTrip()
    table.tripTable()
