# frozen_string_literal: true

require 'spec_helper'

describe 'json output' do
  before :all do
    build_program
  end

  after :all do
    remove_file('csv2table')
  end

  after :each do
    remove_file('output.txt')
  end

  context 'from stdio' do
    context 'with the defaults' do
      it 'creates the json' do
        exec('cat ./spec/json.csv | ./csv2table --json > ./output.txt')

        contents_the_same('data.json')
      end
    end

    context 'delimiter is ,' do
      it 'creates the json' do
        exec('cat ./spec/json.csv | ./csv2table --json -delimit , > ./output.txt')

        contents_the_same('data.json')
      end
    end

    context 'delimiter is tab' do
      it 'creates the json' do
        exec('cat ./spec/json.tsv | ./csv2table --json -delimit "\t" > ./output.txt')

        contents_the_same('data.json')
      end
    end
  end

  context 'as argument' do
    context 'with the defaults' do
      it 'creates the json' do
        exec('./csv2table --json ./spec/json.csv > ./output.txt')

        contents_the_same('data.json')
      end
    end

    context 'delimiter is ,' do
      it 'creates the json' do
        exec('./csv2table -delimit , --json ./spec/json.csv > ./output.txt')

        contents_the_same('data.json')
      end
    end

    context 'delimiter is tab' do
      it 'creates the json' do
        exec('./csv2table -delimit "\t" --json ./spec/json.tsv > ./output.txt')

        contents_the_same('data.json')
      end
    end
  end
end
