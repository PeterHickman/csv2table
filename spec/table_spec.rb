require 'spec_helper'

describe 'convert csv data' do
  before :all do
    build_program
  end

  after :all do
    remove_file('csv2table')
  end

  after :each do
    remove_file('output.txt')
  end

  context 'table output' do
    context 'from stdio' do
      context 'with the defaults' do
        it 'creates the table' do
          s = exec('cat ./spec/data.csv | ./csv2table > ./output.txt')
          expect(s).to eq(0), "csv2table should run without error, got #{s}"

          contents_the_same
        end
      end

      context 'delimiter is ,' do
        it 'creates the table' do
          s = exec('cat ./spec/data.csv | ./csv2table -delimit , > ./output.txt')
          expect(s).to eq(0), "csv2table should run without error, got #{s}"

          contents_the_same
        end
      end

      context 'delimiter is tab' do
        it 'creates the table' do
          s = exec('cat ./spec/data.tsv | ./csv2table -delimit "\t" > ./output.txt')
          expect(s).to eq(0), "csv2table should run without error, got #{s}"

          contents_the_same
        end
      end
    end

    context 'as argument' do
      context 'with the defaults' do
        it 'creates the table' do
          s = exec('./csv2table ./spec/data.csv > ./output.txt')
          expect(s).to eq(0), "csv2table should run without error, got #{s}"

          contents_the_same
        end
      end

      context 'delimiter is ,' do
        it 'creates the table' do
          s = exec('./csv2table -delimit , ./spec/data.csv > ./output.txt')
          expect(s).to eq(0), "csv2table should run without error, got #{s}"

          contents_the_same
        end
      end

      context 'delimiter is tab' do
        it 'creates the table' do
          s = exec('./csv2table -delimit "\t" ./spec/data.tsv > ./output.txt')
          expect(s).to eq(0), "csv2table should run without error, got #{s}"

          contents_the_same
        end
      end
    end
  end
end
