# frozen_string_literal: true

require 'spec_helper'

describe 'rename columns' do
  before :all do
    build_program
  end

  after :all do
    remove_file('csv2')
  end

  after :each do
    remove_file('output.txt')
  end

  context 'for table output' do
    it 'uses the new names' do
      exec('./csv2 --table --names first,second ./spec/data.csv > output.txt')

      contents_the_same('data_names.txt')
    end
  end

  context 'for markdown output' do
    it 'uses the new names' do
      exec('./csv2 --md --names first,second ./spec/data.csv > output.txt')

      contents_the_same('data_names.md')
    end
  end

  context 'for json output' do
    it 'uses the new names' do
      exec('./csv2 --json --names first,second ./spec/data.csv > output.txt')

      contents_the_same('data_names.json')
    end
  end
end
