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

  context 'markdown output' do
    context 'from stdio' do
      context 'with the defaults' do
        it 'creates the table' do
          exec('cat ./spec/data.csv | ./csv2table --md > ./output.txt')

          contents_the_same('data.md')
        end
      end

      context 'delimiter is ,' do
        it 'creates the table' do
          exec('cat ./spec/data.csv | ./csv2table --md -delimit , > ./output.txt')

          contents_the_same('data.md')
        end
      end

      context 'delimiter is tab' do
        it 'creates the table' do
          exec('cat ./spec/data.tsv | ./csv2table --md -delimit "\t" > ./output.txt')

          contents_the_same('data.md')
        end
      end
    end

    context 'as argument' do
      context 'with the defaults' do
        it 'creates the table' do
          exec('./csv2table --md ./spec/data.csv > ./output.txt')

          contents_the_same('data.md')
        end
      end

      context 'delimiter is ,' do
        it 'creates the table' do
          exec('./csv2table -delimit , --md ./spec/data.csv > ./output.txt')

          contents_the_same('data.md')
        end
      end

      context 'delimiter is tab' do
        it 'creates the table' do
          exec('./csv2table -delimit "\t" --md ./spec/data.tsv > ./output.txt')

          contents_the_same('data.md')
        end
      end
    end
  end
end
