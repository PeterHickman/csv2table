# frozen_string_literal: true

require 'English'

def build_program
  system('go build csv2.go')
end

def remove_file(filename)
  system("[ -e #{filename} ] && rm #{filename}")
end

def exec(cmd)
  system(cmd)
  s = $CHILD_STATUS.to_s.split(/\s+/).last.to_i

  expect(s).to eq(0), "should run without error, got #{s}"
end

def contents_the_same(filename)
  expected = File.read("./spec/#{filename}")
  actual   = File.read('./output.txt')

  expect(expected).to eq(actual), 'The output does not match'
end
