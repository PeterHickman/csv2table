def build_program
  system("go build csv2table.go")
end

def remove_file(filename)
  system("[ -e #{filename} ] && rm #{filename}")
end

def exec(cmd)
  system(cmd)
  $?.to_s.split(/\s+/).last.to_i
end

def contents_the_same
  expected = File.read('./spec/data.txt')
  actual   = File.read('./output.txt')

  expect(expected).to eq(actual), 'The output does not match'
end
