def necessary_fuel_for(mass)
  fuel = mass / 3 - 2

  if fuel <= 0
    0
  else
    fuel + necessary_fuel_for(fuel)
  end
end

def mass_without_the_fuel
  puts File.read('input')
    .lines
    .map { |mass| mass.to_i / 3 - 2 }
    .sum(0)
end

def mass_accounting_for_the_fuel
  puts File.read('input')
    .lines
    .map { |mass| necessary_fuel_for mass.to_i }
    .sum(0)
end

mass_accounting_for_the_fuel
