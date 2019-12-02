# frozen_string_literal: true

raw_data = ARGV.first || File.read('input')

memory = raw_data.strip.split(',').map(&:to_i)

def execute(memory, noun, verb)
  memory = memory.dup

  memory[1] = noun
  memory[2] = verb

  address = 0

  loop do
    opcode = memory[address]

    break if opcode == 99

    arg1 = memory[memory[address + 1]]
    arg2 = memory[memory[address + 2]]
    result_index = memory[address + 3]

    memory[result_index] =
      case opcode
      when 1 then arg1 + arg2
      when 2 then arg1 * arg2
      else raise "Unknown opcode #{opcode} at address #{address}"
      end

    address += 4
  end

  memory
end

puts 'Part 1: ', execute(memory, 12, 2).map(&:to_s).join(',')

(0..99).each do |noun|
  (0..99).each do |verb|
    if 19690720 == execute(memory, noun, verb).first
      puts "Part 2: noun=#{noun} verb=#{verb} 100 * noun + verb = #{100 * noun + verb}"
    end
  end
end
