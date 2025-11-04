# frozen_string_literal: true

require 'spec_helper'

describe 'table output' do
  before :all do
    build_program
  end

  after :all do
    remove_file('csv2')
  end

  after :each do
    remove_file('output.txt')
  end

  context 'from stdio' do
    context 'with the defaults' do
      it 'creates the table' do
        exec('cat ./spec/data.csv | ./csv2 > ./output.txt')

        contents_the_same('data.txt')
      end
    end

    context 'delimiter is ,' do
      it 'creates the table' do
        exec('cat ./spec/data.csv | ./csv2 -delimit , > ./output.txt')

        contents_the_same('data.txt')
      end
    end

    context 'delimiter is tab' do
      it 'creates the table' do
        exec('cat ./spec/data.tsv | ./csv2 -delimit "\t" > ./output.txt')

        contents_the_same('data.txt')
      end
    end
  end

  context 'as argument' do
    context 'with the defaults' do
      it 'creates the table' do
        exec('./csv2 ./spec/data.csv > ./output.txt')

        contents_the_same('data.txt')
      end
    end

    context 'delimiter is ,' do
      it 'creates the table' do
        exec('./csv2 -delimit , ./spec/data.csv > ./output.txt')

        contents_the_same('data.txt')
      end
    end

    context 'delimiter is tab' do
      it 'creates the table' do
        exec('./csv2 -delimit "\t" ./spec/data.tsv > ./output.txt')

        contents_the_same('data.txt')
      end
    end
  end
end
