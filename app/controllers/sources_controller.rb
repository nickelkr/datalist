class SourcesController < ApplicationController
  before_action :set_source, only: [:show, :edit, :update, :destroy]
  before_action :authenticate_user!, only: [:edit, :update]
  before_action :is_admin?, only: [:destroy]

  # GET /sources
  # GET /sources.json
  def index
    @sources = Source.paginate(:page => params[:page], :per_page => 10).order('created_at DESC')
  end

  # GET /sources/1
  # GET /sources/1.json
  def show
  end

  # GET /sources/new
  def new
    @source = Source.new
  end

  # GET /sources/1/edit
  def edit
  end

  def search
    @sources = Source.search params[:search]
  end

  # POST /sources
  # POST /sources.json
  def create
    @source = Source.new(source_params)

    respond_to do |format|
      if @source.save
        format.html { redirect_to @source, notice: 'Source was successfully created.' }
        format.json { render :show, status: :created, location: @source }
      else
        format.html { render :new }
        format.json { render json: @source.errors, status: :unprocessable_entity }
      end
    end
  end

  # PATCH/PUT /sources/1
  # PATCH/PUT /sources/1.json
  def update
    respond_to do |format|
      if @source.update(source_params)
        format.html { redirect_to @source, notice: 'Source was successfully updated.' }
        format.json { render :show, status: :ok, location: @source }
      else
        format.html { render :edit }
        format.json { render json: @source.errors, status: :unprocessable_entity }
      end
    end
  end

  # DELETE /sources/1
  # DELETE /sources/1.json
  def destroy
    @source.destroy
    redirect_to(sources_url, notice: 'You need to be an admin to destroy source.') unless is_admin?
    respond_to do |format|
      format.html { redirect_to sources_url, notice: 'Source was successfully destroyed.' }
      format.json { head :no_content }
    end
  end

  private
    # Use callbacks to share common setup or constraints between actions.
    def set_source
      @source = Source.find(params[:id])
    end

    # Checks if is admin
    def is_admin?
      current_user.admin
    end

    # Never trust parameters from the scary internet, only allow the white list through.
    def source_params
      params.require(:source).permit(:name, :url, :description)
    end
end
